function App() {
  const apiUrl = import.meta.env.VITE_API_URL;

  const query = `
    mutation {
      createUser(input: {
        username: "DaGoat"
        password: "jsadhsakjdhas"
        email: "alyan@gmail.com"
      }) {
        id
        email
        username
      }
    }
  `;
  
  const handleClick = async () => {
    const response = await fetch(apiUrl+"/query", {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ query }),
    });
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
