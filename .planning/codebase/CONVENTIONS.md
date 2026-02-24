# Coding Conventions

**Analysis Date:** 2026-02-24

## Naming Patterns

**Files:**
- Vue Components: PascalCase with descriptive names - `BaseButton.vue`, `BaseInputText.vue`
- Composable Functions: camelCase with `use` prefix - `useTheme.ts`
- Store Definitions: camelCase with `use` prefix - `useCounterStore`
- Type Definitions: PascalCase - `types.ts`
- Go Files: snake_case - `handlers.go`, `models.go`, `main.go`

**Functions:**
- TypeScript/Vue: camelCase for regular functions, PascalCase for components
- Go: PascalCase for exported functions, camelCase for package-private

**Variables:**
- TypeScript/Vue: camelCase - `slotContent`, `wrapper`, `themePreference`
- Go: snake_case - `db`, `cfg`, `t` (testing), `w` (httptest.ResponseRecorder)

**Types/Interfaces:**
- TypeScript: PascalCase - `MaybeRef<T>`, `MaybeRefOrGetter<T>`
- Go: PascalCase - `Base`, `HealthCheck` (handler)

**Constants:**
- TypeScript/Vue: UPPER_SNAKE_CASE for configuration constants
- Go: Mixed case based on visibility, typically PascalCase for package-level constants

## Code Style

**Formatting Tool:** Prettier (v2.8.8)

**Settings** (`frontend/.prettierrc.json`):
```json
{
  "semi": false,
  "tabWidth": 2,
  "singleQuote": true,
  "printWidth": 100,
  "trailingComma": "none"
}
```

**Key Style Rules:**
- No semicolons
- 2-space indentation
- Single quotes for strings
- 100 character line width
- No trailing commas

**Styling:**
- SCSS used for component styles
- CSS Modules enabled via `<style lang="scss" module>`
- SCSS mixins and placeholders from `@/design/index.scss`

## Linting

**Tool:** ESLint (v8.57.1)

**Configuration** (`frontend/.eslintrc.cjs`):
```javascript
{
  extends: [
    'plugin:vue/vue3-essential',
    'eslint:recommended',
    '@vue/eslint-config-typescript',
    '@vue/eslint-config-prettier/skip-formatting'
  ],
  ignorePatterns: ['dist/**', 'coverage/**', 'node_modules/**']
}
```

**Backend Linting:**
- `golangci-lint` configured in Makefile
- `go vet` for static analysis

## Import Organization

**TypeScript/Vue:**
- Path aliases: `@/*` maps to `./src/*`
- Auto-imports configured via `unplugin-auto-import` for Vue and VueUse APIs
- Component auto-imports via `unplugin-vue-components`
- ESLint patch (`@rushstack/eslint-patch`) enables modern module resolution

**Standard import order observed:**
1. Third-party packages
2. Path-mapped imports (using `@/`)
3. Relative imports

**Example from `frontend/src/main.ts`:**
```typescript
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from '@/App.vue'
import router from '@/router/index'
```

**Go:**
- Standard library imports first
- Third-party packages second
- Internal package imports third
- Group with blank line between groups

## Error Handling

**TypeScript/Vue:**
- Prop validation using validators: `BaseInputText.vue` uses validator for `type` prop
- Reactive state with Pinia stores
- No explicit try/catch patterns observed in current codebase

**Example from `BaseInputText.vue`:**
```typescript
defineProps({
  type: {
    type: String,
    default: 'text',
    validator: (value: string) =>
      ['email', 'number', 'password', 'search', 'tel', 'text', 'url'].includes(value)
  }
})
```

**Go:**
- Explicit error returns from functions
- Gin handlers return HTTP status codes
- Testing uses testify `assert` for error validation

**Example from `handlers.go`:**
```go
func HealthCheck(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "status":    "ok",
    "timestamp": time.Now().Unix(),
    "service":   "erp-dev-bench",
  })
}
```

## Logging

**Backend (Go):**
- Uses `go.uber.org/zap` (structured logging)
- Not yet implemented in handlers (placeholder code only)

**Frontend:**
- No centralized logging framework detected
- `console` methods for debugging
- VueUse integrations may provide logging utilities

## Comments & Documentation

**TypeScript/Vue:**
- JSDoc comments observed on types in `types.ts`
- Inline comments for complex logic (limited)
- No explicit documentation generation configured

**Example from `types.ts`:**
```typescript
/**
 * It could be a ref, or a plain value
 */
export type MaybeRef<T> = T | Ref<T>
```

**Go:**
- Standard Go doc comments (not extensively used in current code)
- README with API endpoint documentation

## Function & Component Design

**Vue Components:**
- `<script setup lang="ts">` syntax (Composition API)
- Props defined with `defineProps`
- v-model support via `defineModel`
- SCSS Modules for scoped styles

**Example from `BaseButton.spec.ts` and components:**
```vue
<script setup lang="ts">
defineProps({
  // Props defined with TypeScript
})
</script>

<template>
  <!-- Template content -->
</template>

<style lang="scss" module>
/* Scoped styles */
</style>
```

**Pinia Stores:**
- Defined with `defineStore`
- State returns object
- Actions modify state

**Example from `counter.ts`:**
```typescript
export const useCounterStore = defineStore('counter', {
  state: () => ({ count: 0 }),
  actions: {
    increment() { this.count++ }
  }
})
```

## Configuration Conventions

**TypeScript Configuration:**
- Project references system (`tsconfig.json` references)
- Separate configs for app (`tsconfig.app.json`) and vitest (`tsconfig.vitest.json`)
- Path aliases: `@/*` → `./src/*`
- DOM lib enabled for Vue

**Backend Configuration:**
- Viper for configuration management
- Environment-based config loading (`.env.example` present)
- Configuration loaded before server start

## Directory Structure Conventions

**Frontend (`frontend/src/`):**
- Co-located test files (same directory, `.spec.ts` suffix)
- TypeScript source files alongside Vue components
- Clear separation: `components/`, `pages/`, `stores/`, `composables/`, `router/`, `layouts/`, `assets/`, `design/`

**Backend (`backend/`):**
- Go package structure with `internal/` for private packages
- `cmd/` for main entry point
- `tests/` directory with `unit/`, `integration/`, `e2e/` subdirectories
- Each Go package contains its own tests in same directory (convention)
