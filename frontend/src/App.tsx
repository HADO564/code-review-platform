function App() {
  const apiUrl = import.meta.env.VITE_API_URL;

  const query = `
    query {
      user(id: "d024b95c-2fd1-40d1-bdca-0eb3aff31b1d"){
        id
        username
        email
      }
    }
  `;
  
  const handleClick = async () => {
    const response = await fetch(apiUrl+"/query", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer ' + 'd024b95c-2fd1-40d1-bdca-0eb3aff31b1d'
      },
      body: JSON.stringify({ query }),
    });

    if (!response.ok) {
      console.log(response.status)
      return
    }
    const { data } = await response.json();
    console.log(data);
  }
  return (
    <div className='bg-red-600 text-white'>
      <button className='bg-purple-700' onClick={handleClick}>Get the GraphQL Data</button>
    </div>
  )
}

export default App
