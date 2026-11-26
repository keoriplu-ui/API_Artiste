import { useEffect, useState, type FormEvent } from 'react'
import { API_URL_Concert , type Concert } from '../structure/structure.ts'

export function useConcerts() {
  const [concerts, setConcerts] = useState<Concert[]>([])
  const [date, setDate] = useState("")
  const [artistId, setArtistId] = useState<number | null>(null)
  const [errorconcert, setErrorconcert] = useState("")
  const [editingIdconcert, setEditingIdconcert] = useState<number | null>(null)
  const [editDate, setEditDate] = useState("")
  const [editArtistId, setEditArtistId] = useState<number | null>(null)

  async function loadConcerts() {
    try {
      const res = await fetch(`${API_URL_Concert}/concerts`)
      if (!res.ok) throw new Error("Erreur lors du chargement")
      setConcerts(await res.json())
      setErrorconcert("")
    } catch {
      setErrorconcert("Impossible de joindre l'API")
    }
  }

  async function addConcert(e: FormEvent<HTMLFormElement>) {
    e.preventDefault()
    try {
      const res = await fetch(`${API_URL_Concert}/concerts`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ date, artist_id: artistId })
      })
      if (!res.ok) {
        const data = await res.json()
        setErrorconcert(data.error)
        return
      }
      setDate("")
      setArtistId(null)
      loadConcerts()
    } catch {
      setErrorconcert("Impossible de joindre l'API")
    }
  }

  async function deleteConcert(id: number) {
    try {
      await fetch(`${API_URL_Concert}/concerts/${id}`, { method: "DELETE" })
      loadConcerts()
    } catch {
      setErrorconcert("Impossible de joindre l'API")
    }
  }

  async function updateConcert(id: number) {
    try {
      const res = await fetch(`${API_URL_Concert}/concerts/${id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ date: editDate, artist_id: editArtistId }),
      })
      if (!res.ok) {
        const data = await res.json()
        setErrorconcert(data.error)
        return
      }
      setEditingIdconcert(null)
      setErrorconcert("")
      loadConcerts()
    } catch {
      setErrorconcert("Impossible de joindre l'API")
    }
  }

  function startEditconcert(a: Concert) {
    setEditingIdconcert(a.ID)
    setEditDate(a.date)
    setEditArtistId(a.artist_id)
  }

  function cancelEditconcert() {
    setEditingIdconcert(null)
  }

  useEffect(() => {
    loadConcerts()
  }, [])

  return {
    concerts,
    date, setDate,
    artistId, setArtistId,
    errorconcert,
    editingIdconcert,
    editDate, setEditDate,
    editArtistId, setEditArtistId,
    addConcert,
    deleteConcert,
    updateConcert,
    startEditconcert,
    cancelEditconcert,
  }
}