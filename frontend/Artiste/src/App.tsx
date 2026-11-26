import './App.css'
import { useArtists } from './handlersArtists/useArtists.tsx'
import { useConcerts } from './handlersConcerts/useConcerts.tsx'

export default function App() {
  const {
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
  } = useArtists()

  const {
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
  } = useConcerts()

  return (
    <div className="app">
      <h1>Fertival Random</h1>
      <h2>Artistes</h2>

      {error && <p className="error">{error}</p>}

      <form className="modif" onSubmit={addArtist}>
        <input
          placeholder="Nom"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <input
          placeholder="Genre"
          value={genre}
          onChange={(e) => setGenre(e.target.value)}
        />
        <button type="submit">Ajouter</button>
      </form>

      <ul>
        {artists.map((a) => (
          <li key={a.ID}>
            {editingId === a.ID ? (
              <>
                <input
                  value={editName}
                  onChange={(e) => setEditName(e.target.value)}
                />
                <input
                  value={editGenre}
                  onChange={(e) => setEditGenre(e.target.value)}
                />
                <button onClick={() => updateArtist(a.ID)}>Enregistrer</button>
                <button onClick={cancelEdit}>Annuler</button>
              </>
            ) : (
              <>
                <span>
                  <strong>{a.name}</strong> ({a.genre})
                </span>
                <button onClick={() => startEdit(a)}>Modifier</button>
                <button onClick={() => deleteArtist(a.ID)}>Supprimer</button>
              </>
            )}
          </li>
        ))}
      </ul>

      <h2>Concerts</h2>
      {errorconcert && <p className="error">{errorconcert}</p>}

      <form className="modif" onSubmit={addConcert}>
        <input
          placeholder="Date"
          value={date}
          onChange={(e) => setDate(e.target.value)}
        />
        <input
          placeholder="Artiste ID"
          value={artistId ?? ''}
          onChange={(e) =>
            setArtistId(e.target.value === '' ? null : Number(e.target.value))
          }
        />
        <button type="submit">Ajouter</button>
      </form>

      <ul>
        {concerts.map((c) => (
          <li key={c.ID}>
            {editingIdconcert === c.ID ? (
              <>
                <input
                  value={editDate}
                  onChange={(e) => setEditDate(e.target.value)}
                />
                <input
                  value={editArtistId ?? ''}
                  onChange={(e) =>
                    setEditArtistId(e.target.value === '' ? null : Number(e.target.value))
                  }
                />
                <button onClick={() => updateConcert(c.ID)}>Enregistrer</button>
                <button onClick={cancelEditconcert}>Annuler</button>
              </>
            ) : (
              <>
                <span>
                  <strong>{c.date}</strong> ({c.artist_id})
                </span>
                <button onClick={() => startEditconcert(c)}>Modifier</button>
                <button onClick={() => deleteConcert(c.ID)}>Supprimer</button>
              </>
            )}
          </li>
        ))}
      </ul>
    </div>
  )
}