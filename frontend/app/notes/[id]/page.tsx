'use client'

import { useEffect, useState } from 'react'
import { useRouter, useParams } from 'next/navigation'
import Link from 'next/link'
import { getNote, deleteNote, uploadImage, isAuthenticated } from '@/lib/api'

interface Note {
  id: number
  title: string
  content: string
  image_url?: string
  created_at: string
}

export default function NoteDetail() {
  const router = useRouter()
  const params = useParams()
  const noteId = params.id as string

  const [note, setNote] = useState<Note | null>(null)
  const [loading, setLoading] = useState(true)
  const [uploading, setUploading] = useState(false)
  const [error, setError] = useState('')

  useEffect(() => {
    if (!isAuthenticated()) {
      router.push('/login')
      return
    }

    loadNote()
  }, [noteId, router])

  const loadNote = async () => {
    try {
      const data = await getNote(noteId)
      setNote(data)
    } catch (error) {
      setError('Failed to load note')
    } finally {
      setLoading(false)
    }
  }

  const handleDelete = async () => {
    if (!confirm('Are you sure you want to delete this note?')) return

    try {
      await deleteNote(noteId)
      router.push('/notes')
    } catch (error) {
      alert('Failed to delete note')
    }
  }

  const handleImageUpload = async (e: React.ChangeEvent<HTMLInputElement>) => {
    const file = e.target.files?.[0]
    if (!file) return

    setUploading(true)
    try {
      const updatedNote = await uploadImage(noteId, file)
      setNote(updatedNote)
    } catch (error: any) {
      alert(error.message || 'Failed to upload image')
    } finally {
      setUploading(false)
    }
  }

  if (loading) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
          <p className="mt-4 text-gray-600">Loading note...</p>
        </div>
      </div>
    )
  }

  if (error || !note) {
    return (
      <div className="min-h-screen bg-gray-50 flex items-center justify-center">
        <div className="text-center">
          <h2 className="text-2xl font-semibold text-gray-700 mb-4">Note not found</h2>
          <Link
            href="/notes"
            className="bg-blue-600 text-white px-6 py-2 rounded-lg font-semibold hover:bg-blue-700 transition inline-block"
          >
            Back to Notes
          </Link>
        </div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-4xl mx-auto px-4 py-4 sm:px-6 lg:px-8">
          <div className="flex items-center justify-between">
            <Link
              href="/notes"
              className="text-gray-600 hover:text-gray-900 transition"
            >
              ← Back to Notes
            </Link>
            <button
              onClick={handleDelete}
              className="bg-red-600 text-white px-4 py-2 rounded-lg font-semibold hover:bg-red-700 transition"
            >
              Delete Note
            </button>
          </div>
        </div>
      </header>

      {/* Content */}
      <main className="max-w-4xl mx-auto px-4 py-8 sm:px-6 lg:px-8">
        <div className="bg-white rounded-lg shadow p-8">
          <h1 className="text-3xl font-bold text-gray-900 mb-4">{note.title}</h1>
          
          <p className="text-sm text-gray-500 mb-6">
            Created on {new Date(note.created_at).toLocaleDateString()} at{' '}
            {new Date(note.created_at).toLocaleTimeString()}
          </p>

          {note.image_url && (
            <div className="mb-6">
              <img
                src={`http://localhost:8080${note.image_url}`}
                alt={note.title}
                className="w-full max-h-96 object-contain rounded-lg"
              />
            </div>
          )}

          <div className="prose max-w-none mb-6">
            <p className="text-gray-700 whitespace-pre-wrap">{note.content}</p>
          </div>

          {/* Image Upload */}
          <div className="border-t pt-6">
            <h3 className="text-lg font-semibold mb-3">Add Image</h3>
            <div className="flex items-center space-x-4">
              <input
                type="file"
                accept="image/*"
                onChange={handleImageUpload}
                disabled={uploading}
                className="block w-full text-sm text-gray-500
                  file:mr-4 file:py-2 file:px-4
                  file:rounded-lg file:border-0
                  file:text-sm file:font-semibold
                  file:bg-blue-50 file:text-blue-700
                  hover:file:bg-blue-100
                  disabled:opacity-50 disabled:cursor-not-allowed"
              />
              {uploading && (
                <div className="text-sm text-gray-600">Uploading...</div>
              )}
            </div>
          </div>
        </div>
      </main>
    </div>
  )
}