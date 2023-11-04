import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';

const Home = () => {
  const [usersCount, setUsersCount] = useState(0);
  const [micropostsCount, setMicropostsCount] = useState(0);

  useEffect(() => {
    axios.get('http://localhost:8080/api/users/count')
      .then(response => {
        setUsersCount(response.data.count); 
      })
      .catch(error => {
        console.error('ユーザー数の取得中にエラーが発生しました：', error);
      });

    axios.get('http://localhost:8080/api/microposts/count')
      .then(response => {
        setMicropostsCount(response.data.count);
      })
      .catch(error => {
        console.error('マイクロポスト数の取得中にエラーが発生しました：', error);
      });
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <p>ユーザー数：{usersCount}</p>
        <p>マイクロポスト数：{micropostsCount}</p>
        <nav>
          <ul>
            <li><Link to="/users">ユーザー一覧</Link></li>
            <li><Link to="/microposts">マイクロポスト一覧</Link></li>
          </ul>
        </nav>
      </header>
    </div>
  );
}

export default Home;
