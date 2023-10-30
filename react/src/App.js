// App.js
import React, { useState, useEffect } from 'react';
import axios from 'axios'; // axios を使用して HTTP リクエストを行う

function App() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    axios.get('http://localhost:8080/api/hello')
      .then(response => {
        setMessage(response.data.message); // レスポンスのメッセージを状態にセット
      })
      .catch(error => {
        console.error('There was an error fetching data:', error);
      });
  }, []); // 空の依存配列を渡すことで、このエフェクトはコンポーネントのマウント時にのみ実行されます

  return (
    <div className="App">
      <header className="App-header">
        <p>
          {message} {/* Go サーバーからのメッセージを表示 */}
        </p>
      </header>
    </div>
  );
}

export default App;
