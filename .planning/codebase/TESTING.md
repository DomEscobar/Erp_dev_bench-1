# Testing Patterns

**Analysis Date:** 2026-02-24

## Test Frameworks

**Frontend (TypeScript/Vue):**
- **Runner:** Vitest (^1.4.0)
- **Assertion:** Built-in `expect` from Vitest
- **Component Testing:** `@vue/test-utils` (^2.4.6)
- **E2E Testing:** Playwright (^1.44.0)
- **Environment:** jsdom for unit tests

**Backend (Go):**
- **Runner:** Go's built-in `testing` package
- **Assertion:** `testify/assert` (v1.11.1)
- **HTTP Testing:** `net/http/httptest`
- **Database Mocking:** `go-sqlmock` (indicated by README)
- **Integration:** Real database testing with SQLite in-memory

## Test File Organization

**Frontend:**
- Co-located test files in same directory as source
- Naming: `*.spec.ts` or `*.test.ts` (both observed)
- Test files alongside component files

**Example:**
```
frontend/src/components/
├── BaseButton.vue
├── BaseButton.spec.ts
├── BaseInputText.vue
└── BaseInputText.vue
```

**Backend (Go):**
- Separate `backend/tests/` directory with:
  - `unit/` - Unit tests for individual packages
  - `integration/` - Integration tests (database, external services)
  - `e2e/` - End-to-end tests (full system)
- Test files named `*_test.go` in package directories or test subdirectories

**Example from backend/tests:**
```
backend/tests/
├── unit/
│   └── handlers_test.go
├── integration/
│   └── integration_test.go
└── e2e/
    └── e2e_test.go
```

## Test Structure

**Vitest Pattern:**
```typescript
import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import Component from './Component.vue'

describe('Component Name', () => {
  it('renders properly', () => {
    const wrapper = mount(Component, {
      props: { /* props */ }
    })
    expect(wrapper.find('selector').exists()).toBe(true)
  })

  it('has expected content', () => {
    // assertions
  })
})
```

**Pattern from `BaseInputText.spec.ts`:**
```typescript
describe('BaseInputText', () => {
  it('renders properly', () => {
    const wrapper = mount(BaseInputText, {
      props: { modelValue: '' }
    })
    expect(wrapper.find('input').exists()).toBe(true)
  })
})
```

**Go Testing Pattern:**
```go
package handlers

import (
  "testing"
  "net/http/httptest"
  "github.com/stretchr/testify/assert"
  "github.com/gin-gonic/gin"
)

func TestHealthCheck(t *testing.T) {
  gin.SetMode(gin.TestMode)
  
  w := httptest.NewRecorder()
  c, _ := gin.CreateTestContext(w)
  
  handlers.HealthCheck(c)
  
  assert.Equal(t, http.StatusOK, w.Code)
  assert.Contains(t, w.Body.String(), "ok")
}
```

**Go Integration Pattern:**
```go
package integration

import (
  "testing"
  "github.com/stretchr/testify/assert"
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

func TestDatabaseConnection(t *testing.T) {
  db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
  assert.NoError(t, err)
  assert.NotNil(t, db)
}
```

## Mocking

**Frontend:**
- Component mocking via `@vue/test-utils` (`mount`, `shallowMount`)
- Props passed as options
- Slots mocked via `slots` option
- No explicit mocking library for functions/stubs

**Examples from `BaseButton.spec.ts`:**
```typescript
// shallowMount isolates component child dependencies
const { element } = shallowMount(BaseButton, {
  slots: {
    default: slotContent
  }
})
```

**Backend (Go):**
- HTTP testing with `httptest.NewRecorder()` and `gin.CreateTestContext()`
- Database mocking with in-memory SQLite for integration tests
- `go-sqlmock` mentioned in README (not in sample files)
- Gin test mode to avoid real HTTP server

**Example mocking pattern:**
```go
w := httptest.NewRecorder()
c, _ := gin.CreateTestContext(w)
handlers.HealthCheck(c)
assert.Equal(t, http.StatusOK, w.Code)
```

## Test Commands

**Frontend (`frontend/package.json` scripts):**
```bash
npm run test:unit     # Run vitest
npm run test:e2e      # Run Playwright tests
npm run test:e2e:ui   # Playwright with UI
npm run type-check    # Vue TypeScript checking
```

**Backend (`backend/Makefile` targets):**
```bash
make test             # Run all tests with coverage
make coverage         # Generate HTML coverage report
go test ./... -v      # Verbose test run
go test ./tests/unit/... -v  # Unit tests only
go test ./tests/integration/... -v  # Integration tests only
go test ./tests/e2e/... -v  # E2E tests only
```

## Coverage

**Backend:**
- Coverage enabled via `go test -cover -coverprofile=coverage.out`
- HTML report generation: `go tool cover -html=coverage.out -o coverage.html`
- Integration with Makefile: `make coverage`

**Frontend:**
- Vitest config doesn't explicitly show coverage configuration
- Coverage likely available via `vitest --coverage`
- Coverage directory pattern (if following Vitest defaults): `coverage/`

**Ignored for Coverage:**
- Frontend: `dist/**`, `coverage/**`, `node_modules/**` in ESLint config
- Type test files excluded: `src/**/__tests__/*` in `tsconfig.app.json`

## Test Types

**Frontend:**
1. **Unit Tests:** Component behavior and rendering with Vitest + Vue Test Utils
2. **E2E Tests:** Full application flow with Playwright (accessing `/` route)

**Example E2E test (`frontend/e2e/vue.spec.ts`):**
```typescript
import { test, expect } from '@playwright/test'

test('visits the app root url', async ({ page }) => {
  await page.goto('/')
  await expect(page.locator('h1')).toHaveText('Home Page')
})
```

**Backend (Go):**
1. **Unit Tests:** Individual handler functions using httptest
2. **Integration Tests:** Database connectivity, GORM operations with in-memory SQLite
3. **E2E Tests:** Placeholder test indicating framework readiness (`TestE2EPlaceholder`)

## Test Setup Patterns

**Vitest Configuration** (`frontend/vitest.config.ts`):
```typescript
export default mergeConfig(
  viteConfig,
  defineConfig({
    test: {
      environment: 'jsdom',
      exclude: [...configDefaults.exclude, 'e2e/*'],
      root: fileURLToPath(new URL('./', import.meta.url)),
      transformMode: {
        web: [/\.[jt]sx$/],
      },
    }
  })
)
```

**Go Testing Environment:**
- Test files use Go's standard `testing` package
- Table-driven tests pattern (not observed but idiomatic Go)
- Setup/teardown via `t.Cleanup()` or explicit functions

## Common Patterns

**Assertions:**
- Frontend: `expect(value).toBe(expected)`, `expect(element).toExist()`
- Backend: `assert.Equal(t, actual, expected)`, `assert.NoError(t, err)`, `assert.NotNil(t, value)`

**Async Testing:**
- Frontend: `async/await` with Playwright for async operations
- Go: Synchronous by nature, goroutines tested with channels or `t.Parallel()`

**Error Testing:**
- Frontend: `expect(fn).toThrow()`, checking error messages
- Backend: `assert.Error(t, err)`, `assert.Contains(t, err.Error(), "expected")`

## Fixtures & Test Data

**Frontend:**
- Inline test data within test files
- No dedicated fixtures directory observed
- Mock props and slots directly in test setup

**Backend:**
- In-memory SQLite database: `file::memory:?cache=shared`
- Test data created inline in test functions
- No external fixture files observed

## Observability in Tests

**Console Output:**
- Vitest: Console output captured, can use `vi.spyOn(console, 'log')` for verification
- Go: `t.Log()` for test logging

**Debugging:**
- Frontend: `debugger` statements or VS Code breakpoints with Vitest
- Backend: Delve debugger or `fmt.Println`/`t.Log` for quick inspection

## Test Dependencies

**Frontend critical test packages:**
- `vitest` (^1.4.0)
- `@vue/test-utils` (^2.4.6)
- `@playwright/test` (^1.44.0)
- `jsdom` (^22.0.0) - DOM environment

**Backend critical test packages:**
- `github.com/stretchr/testify` (v1.11.1)
- `golang.org/x/net` for HTTP testing (standard library also sufficient)
- `go-sqlmock` (mentioned, not seen in sample files)

## Recommendations

**Frontend:**
- Keep test files co-located with components for visibility
- Use `shallowMount` for isolation, `mount` for full rendering
- Test both happy paths and edge cases
- Add error state tests where applicable

**Backend:**
- Place unit tests in same package directory as source (standard Go)
- Use table-driven tests for multiple input/output scenarios
- Mock external dependencies (databases, APIs) with in-memory or mock implementations
- Keep e2e tests minimal and focused on critical user journeys
