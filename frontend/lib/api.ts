// API utility functions for interacting with backend

const API_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

// Get token from localStorage
const getToken = (): string | null => {
  if (typeof window !== 'undefined') {
    return localStorage.getItem('token');
  }
  return null;
};

// Register user
export const register = async (name: string, email: string, password: string) => {
  const response = await fetch(`${API_URL}/api/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name, email, password }),
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Registration failed');
  }

  return response.json();
};

// Login user
export const login = async (email: string, password: string) => {
  const response = await fetch(`${API_URL}/api/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Login failed');
  }

  const data = await response.json();
  
  // Store token in localStorage
  if (data.token) {
    localStorage.setItem('token', data.token);
    localStorage.setItem('user', JSON.stringify(data.user));
  }

  return data;
};

// Logout user
export const logout = () => {
  localStorage.removeItem('token');
  localStorage.removeItem('user');
};

// Get all notes
export const getNotes = async () => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const response = await fetch(`${API_URL}/api/notes`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error('Failed to fetch notes');
  }

  return response.json();
};

// Get single note
export const getNote = async (id: string) => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const response = await fetch(`${API_URL}/api/notes/${id}`, {
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error('Failed to fetch note');
  }

  return response.json();
};

// Create note
export const createNote = async (title: string, content: string) => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const response = await fetch(`${API_URL}/api/notes`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify({ title, content }),
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to create note');
  }

  return response.json();
};

// Update note
export const updateNote = async (id: string, title: string, content: string) => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const response = await fetch(`${API_URL}/api/notes/${id}`, {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `Bearer ${token}`,
    },
    body: JSON.stringify({ title, content }),
  });

  if (!response.ok) {
    throw new Error('Failed to update note');
  }

  return response.json();
};

// Delete note
export const deleteNote = async (id: string) => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const response = await fetch(`${API_URL}/api/notes/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
  });

  if (!response.ok) {
    throw new Error('Failed to delete note');
  }

  return response.json();
};

// Upload image for note
export const uploadImage = async (noteId: string, file: File) => {
  const token = getToken();
  if (!token) throw new Error('No token found');

  const formData = new FormData();
  formData.append('image', file);

  const response = await fetch(`${API_URL}/api/notes/${noteId}/upload`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${token}`,
    },
    body: formData,
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'Failed to upload image');
  }

  return response.json();
};

// Check if user is authenticated
export const isAuthenticated = (): boolean => {
  return getToken() !== null;
};

// Get current user from localStorage
export const getCurrentUser = () => {
  if (typeof window !== 'undefined') {
    const userStr = localStorage.getItem('user');
    return userStr ? JSON.parse(userStr) : null;
  }
  return null;
};