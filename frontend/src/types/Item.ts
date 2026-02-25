export interface Item {
  id: number
  name: string
  description: string
  price: number
  sku: string
  categoryId: number
  createdAt: string
  updatedAt: string
  deletedAt?: string
}
