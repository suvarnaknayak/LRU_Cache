import React, { useState, useEffect } from 'react';
import { initWebSocket, sendWebSocketMessage } from '../ws/websocket';

const CacheManager = () => {
  const [key, setKey] = useState('');
  const [value, setValue] = useState('');
  const [expiration, setExpiration] = useState(0);
  const [cache, setCache] = useState([]);

  useEffect(() => {
    const ws = initWebSocket((event) => {
      const data = JSON.parse(event.data);
      setCache(data);
    });

    return () => {
      ws.close(); // Clean up the WebSocket connection
    };
  }, []); // Empty dependency array ensures this effect runs once on mount

  const handleSet = async () => {
    await fetch('http://localhost:8080/cache/set', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ key, value, expiration }),
    });
    sendWebSocketMessage({ type: 'update' });
  };

  return (
    <div>
      <h1>LRU Cache Manager</h1>
      <input
        type="text"
        placeholder="Key"
        value={key}
        onChange={(e) => setKey(e.target.value)}
      />
      <input
        type="text"
        placeholder="Value"
        value={value}
        onChange={(e) => setValue(e.target.value)}
      />
      <input
        type="number"
        placeholder="Expiration (seconds)"
        value={expiration}
        onChange={(e) => setExpiration(Number(e.target.value))}
      />
      <button onClick={handleSet}>Set</button>
      <ul>
        {cache.map((item, index) => (
          <li key={index}>
            {item.key}: {item.value} (expires in {item.expiration}s)
          </li>
        ))}
      </ul>
    </div>
  );
};

export default CacheManager;
