import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import LoadingIndicator from './LoadingIndicator';

const Home = () => {
  const [usersCount, setUsersCount] = useState(0);
  const [micropostsCount, setMicropostsCount] = useState(0);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    Promise.all([
      axios.get('http://localhost:8080/api/users/count'),
      axios.get('http://localhost:8080/api/microposts/count')
    ])
    .then(([usersResponse, micropostsResponse]) => {
      setUsersCount(usersResponse.data.count);
      setMicropostsCount(micropostsResponse.data.count);
    })
    .catch(error => {
      console.error('データの取得中にエラーが発生しました：', error);
    })
    .finally(() => {
      setLoading(false);
    });
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <p>ユーザー数：{loading ? <LoadingIndicator /> : usersCount}</p>
        <p>マイクロポスト数：{loading ? <LoadingIndicator /> : micropostsCount}</p>
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
