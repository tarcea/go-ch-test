import React, { useEffect, useState } from 'react';
import './App.css';

function App() {
  const [games, setGame] = useState([]);

  const fetchGames = async () => {
    const games = await fetch('http://localhost:8080/games');
    const data = await games.json();
    setGame(data);
  };

  useEffect(() => {
    fetchGames();
  }, []);
  console.log(games);
  return (
    <div className='App'>
      {games.length && (
        <div>
          {games[0].turn ? "white's turn" : "black's turn"}
          <section>
            {games[0].board?.map((line, j) => {
              // return line.map((l, i) => {
              return (
                <div className='line' key={j}>
                  {line.map((l, i) => (
                    <p className='pl' key={i + 99}>
                      {l}
                    </p>
                  ))}
                </div>
              );
              // });
            })}
          </section>
        </div>
      )}
    </div>
  );
}

export default App;
