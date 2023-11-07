import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import LoadingIndicator from './LoadingIndicator';
import { API_BASE_URL, DEFAULT_TIMEOUT } from '../config';

const fetchData = async (endpoint) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/${endpoint}/count`, {
      timeout: DEFAULT_TIMEOUT,
    });
    return { count: response.data.count };
  } catch (error) {
    if (error.code === 'ECONNABORTED') {
      return { error: 'timeout' };
    } else {
      return { error: 'error', message: error.message };
    }
  }
};

const Home = () => {
  const [usersCount, setUsersCount] = useState(null);
  const [micropostsCount, setMicropostsCount] = useState(null);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    const loadCounts = async () => {
      try {
        const usersPromise = fetchData('users');
        const micropostsPromise = fetchData('microposts');
        const minimumLoadTimePromise = new Promise(resolve => setTimeout(resolve, 500));
        const promises = [usersPromise, micropostsPromise, minimumLoadTimePromise];

        const [users, microposts] = await Promise.all(promises);
        console.log(users)

        setUsersCount(users.count);
        setMicropostsCount(microposts.count);

        if (users.error || microposts.error) {
          setErrorMessage(users.error ?? microposts.error)
        } else {
          setUsersCount(users.count);
          setMicropostsCount(microposts.count);
        }
      } finally {
        setLoading(false);
      }
    };

    loadCounts();
  }, []);

  const displayContent = (content) => {
    if (loading) return <LoadingIndicator />;
    if (errorMessage) return <span>Error: {errorMessage}</span>;
    return content ?? '0';
  };

  return (
    <div className="App">
      <header className="App-header">
        <p>ユーザー数：{displayContent(usersCount)}</p>
        <p>マイクロポスト数：{displayContent(micropostsCount)}</p>
        <nav>
          <ul>
            <li><Link to="/users">ユーザー一覧</Link></li>
            <li><Link to="/microposts">マイクロポスト一覧</Link></li>
          </ul>
        </nav>
      </header>
    </div>
  );
};

export default Home;
