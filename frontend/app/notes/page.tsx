'use client'

import { useEffect, useState } from 'react'
import { useRouter } from 'next/navigation'
import Link from 'next/link'
import { getNotes, logout, isAuthenticated, getCurrentUser } from '@/lib/api'

interface Note {
  id: number
  title: string
  content: string
  image_url?: string
  created_at: string
}

export default function Notes() {
  const router = useRouter()
  const [notes, setNotes] = useState<Note[]>([])
  const [loading, setLoading] = useState(true)
  const [user, setUser] = useState<any>(null)

  useEffect(() => {
    // Check authentication
    if (!isAuthenticated()) {
      router.push('/login')
      return
    }

    setUser(getCurrentUser())
    loadNotes()
  }, [router])

  const loadNotes = async () => {
    try {
      const data = await getNotes()
      setNotes(data.notes || [])
    } catch (error) {
      console.error('Failed to load notes:', error)
    } finally {
      setLoading(false)
    }
  }

  const handleLogout = () => {
    logout()
    router.push('/login')
  }

  return (
    <div className="min-h-screen bg-gray-50">
      {/* Header */}
      <header className="bg-white shadow-sm">
        <div className="max-w-7xl mx-auto px-4 py-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center">
            <div>
              <h1 className="text-2xl font-bold text-gray-900">📝 My Notes</h1>
              {user && <p className="text-sm text-gray-600">Welcome, {user.name}!</p>}
            </div>
            <div className="space-x-4">
              <Link
                href="/notes/new"
                className="bg-blue-600 text-white px-4 py-2 rounded-lg font-semibold hover:bg-blue-700 transition inline-block"
              >
                + New Note
              </Link>
              <button
                onClick={handleLogout}
                className="bg-gray-200 text-gray-700 px-4 py-2 rounded-lg font-semibold hover:bg-gray-300 transition"
              >
                Logout
              </button>
            </div>
          </div>
        </div>
      </header>

      {/* Content */}
      <main className="max-w-7xl mx-auto px-4 py-8 sm:px-6 lg:px-8">
        {loading ? (
          <div className="text-center py-12">
            <div className="inline-block animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            <p className="mt-4 text-gray-600">Loading notes...</p>
          </div>
        ) : notes.length === 0 ? (
          <div className="text-center py-12">
            <div className="text-6xl mb-4">📝</div>
            <h2 className="text-2xl font-semibold text-gray-700 mb-2">No notes yet</h2>
            <p className="text-gray-600 mb-6">Create your first note to get started!</p>
            <Link
              href="/notes/new"
              className="bg-blue-600 text-white px-6 py-3 rounded-lg font-semibold hover:bg-blue-700 transition inline-block"
            >
              Create Note
            </Link>
          </div>
        ) : (
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {notes.map((note) => (
              <Link
                key={note.id}
                href={`/notes/${note.id}`}
                className="bg-white rounded-lg shadow hover:shadow-lg transition p-6 cursor-pointer"
              >
                <h3 className="text-xl font-semibold text-gray-900 mb-2 truncate">
                  {note.title}
                </h3>
                <p className="text-gray-600 mb-4 line-clamp-3">{note.content}</p>
                {note.image_url && (
                  <div className="mb-4">
                    <img
                      src={`http://localhost:8080${note.image_url}`}
                      alt={note.title}
                      className="w-full h-32 object-cover rounded"
                    />
                  </div>
                )}
                <p className="text-sm text-gray-500">
                  {new Date(note.created_at).toLocaleDateString()}
                </p>
              </Link>
            ))}
          </div>
        )}
      </main>
    </div>
  )
}