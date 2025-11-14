'use client'

import { useEffect } from 'react'
import { useRouter } from 'next/navigation'
import { isAuthenticated } from '@/lib/api'
import Link from 'next/link'

export default function Home() {
  const router = useRouter()

  useEffect(() => {
    // Redirect to notes if already authenticated
    if (isAuthenticated()) {
      router.push('/notes')
    }
  }, [router])

  return (
    <div className="min-h-screen flex items-center justify-center bg-gradient-to-br from-blue-500 to-purple-600">
      <div className="text-center text-white">
        <h1 className="text-6xl font-bold mb-4">📝 Notes Sharing App</h1>
        <p className="text-xl mb-8">Create, share, and manage your notes with ease</p>
        
        <div className="space-x-4">
          <Link 
            href="/login"
            className="bg-white text-blue-600 px-8 py-3 rounded-lg font-semibold hover:bg-gray-100 transition inline-block"
          >
            Login
          </Link>
          <Link 
            href="/register"
            className="bg-transparent border-2 border-white text-white px-8 py-3 rounded-lg font-semibold hover:bg-white hover:text-blue-600 transition inline-block"
          >
            Register
          </Link>
        </div>

        <div className="mt-12 grid grid-cols-3 gap-8 max-w-3xl mx-auto">
          <div className="bg-white/10 backdrop-blur-sm p-6 rounded-lg">
            <div className="text-4xl mb-2">🔒</div>
            <h3 className="font-semibold mb-2">Secure</h3>
            <p className="text-sm">Your notes are protected with JWT authentication</p>
          </div>
          <div className="bg-white/10 backdrop-blur-sm p-6 rounded-lg">
            <div className="text-4xl mb-2">⚡</div>
            <h3 className="font-semibold mb-2">Fast</h3>
            <p className="text-sm">Built with Next.js and Go for blazing speed</p>
          </div>
          <div className="bg-white/10 backdrop-blur-sm p-6 rounded-lg">
            <div className="text-4xl mb-2">📱</div>
            <h3 className="font-semibold mb-2">Responsive</h3>
            <p className="text-sm">Works perfectly on all your devices</p>
          </div>
        </div>
      </div>
    </div>
  )
}