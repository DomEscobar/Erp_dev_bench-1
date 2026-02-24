---
created: 2026-02-24T16:39:58.278Z
title: Items CRUD API + Frontend
area: api
files:
  - backend/internal/handlers/items.go
  - frontend/src/pages/ItemsPage.vue
  - frontend/src/stores/items.ts
---

## Problem

Create a complete Items management system with CRUD API endpoints and frontend table view with add/edit functionality.

## Solution

- Backend (Go): Implement handlers in `backend/internal/handlers/items.go` with GORM model including fields: ID, Name, Description, Price, SKU (unique), CategoryID. Endpoints: GET /api/v1/items, GET /api/v1/items/:id, POST /api/v1/items, PUT /api/v1/items/:id, DELETE /api/v1/items/:id
- Frontend (Vue): Create ItemsPage.vue with ItemsTable.vue and ItemForm.vue components, plus Pinia store items.ts
- Tests: Backend CRUD operations, frontend table rendering and form submission
- KPIs: TypeScript, lint, build, tests passing
- Estimated: 8 files, 15-30 min
