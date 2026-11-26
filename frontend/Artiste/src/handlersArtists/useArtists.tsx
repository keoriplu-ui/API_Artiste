import { useEffect, useState, type FormEvent } from 'react'
import { API_URL_Artist , type Artist } from '../structure/structure.ts'

export function useArtists() {
  const [artists, setArtists] = useState<Artist[]>([])
  const [name, setName] = useState("")
  const [genre, setGenre] = useState("")
  const [error, setError] = useState("")
  const [editingId, setEditingId] = useState<number | null>(null)
  const [editName, setEditName] = useState("")
  const [editGenre, setEditGenre] = useState("")

  async function loadArtists() {
    try {
      const res = await fetch(`${API_URL_Artist}/artists`)
      if (!res.ok) throw new Error("Erreur lors du chargement")
      setArtists(await res.json())
      setError("")
    } catch {
      setError("Impossible de joindre l'API")
    }
  }

  async function addArtist(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    try {
      const res = await fetch(`${API_URL_Artist}/artists`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name, genre }),
      })
      if (!res.ok) {
        const data = await res.json()
        setError(data.erreur)
        return
      }
      setName("")
      setGenre("")
      loadArtists()
    } catch {
      setError("Impossible de joindre l'API")
    }
  }

  async function deleteArtist(id: number) {
    try {
      await fetch(`${API_URL_Artist}/artists/${id}`, { method: "DELETE" })
      loadArtists()
    } catch {
      setError("Impossible de joindre l'API")
    }
  }

  async function updateArtist(id: number) {
    try {
      const res = await fetch(`${API_URL_Artist}/artists/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ name: editName, genre: editGenre }),
      })
      if (!res.ok) {
        const data = await res.json()
        setError(data.error)
        return
      }
      setEditingId(null)
      setError("")
      loadArtists()
    } catch {
      setError("Impossible de joindre l'API")
    }
  }

  function startEdit(a: Artist) {
    setEditingId(a.ID)
    setEditName(a.name)
    setEditGenre(a.genre)
  }

  function cancelEdit() {
    setEditingId(null)
  }

  useEffect(() => {
    loadArtists()
  }, [])

  return {
    artists,
    name, setName,
    genre, setGenre,
    error,
    editingId,
    editName, setEditName,
    editGenre, setEditGenre,
    addArtist,
    deleteArtist,
    updateArtist,
    startEdit,
    cancelEdit,
  }
}