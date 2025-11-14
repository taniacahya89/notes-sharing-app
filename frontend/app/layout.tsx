import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
  title: 'Notes Sharing App',
  description: 'Share your notes with ease',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  )
}