import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { Link } from 'react-router-dom';
import LoadingIndicator from './LoadingIndicator';
import { API_BASE_URL, DEFAULT_TIMEOUT } from '../config';

const fetchData = (endpoint) =>
  axios.get(`${API_BASE_URL}/${endpoint}/count`, { timeout: DEFAULT_TIMEOUT });

const Home = () => {
  const [usersCount, setUsersCount] = useState(null);
  const [micropostsCount, setMicropostsCount] = useState(null);
  const [loading, setLoading] = useState(true);
  const [errorMessage, setErrorMessage] = useState('');

  useEffect(() => {
    const promises = [
      fetchData('users').catch(handleError),
      fetchData('microposts').catch(handleError),
      new Promise((resolve) => setTimeout(resolve, 500)),
    ];

    Promise.all(promises).then(([usersResponse, micropostsResponse]) => {
      if (!errorMessage) {
        setUsersCount(usersResponse?.data?.count);
        setMicropostsCount(micropostsResponse?.data?.count);
      }
    })
    .finally(() => setLoading(false));
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  function handleError(error) {
    const message = error.code === 'ECONNABORTED' ? 'timeout' : 'error';
    setErrorMessage(message);
    return { data: {} };
  }

  const displayContent = (content) => {
    if (loading) return <LoadingIndicator />;
    if (errorMessage) return errorMessage;
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
