import React from 'react';
import { render, screen, within } from '@testing-library/react';
import '@testing-library/jest-dom';
import axios from 'axios';
import Home from './Home';
import { MemoryRouter } from 'react-router-dom';
import { API_BASE_URL } from '../config';
import { act } from 'react-dom/test-utils';

jest.mock('axios');

describe('Home コンポーネント', () => {
  test('ユーザー数とマイクロポスト数が表示される', async () => {
    jest.useFakeTimers();
  
    axios.get.mockImplementation(url => {
      switch (url) {
        case `${API_BASE_URL}/users/count`:
          return Promise.resolve({ data: { count: 10 } });
        case `${API_BASE_URL}/microposts/count`:
          return Promise.resolve({ data: { count: 5 } });
        default:
          throw new Error('not found');
      }
    });
  
    await act(async () => {
      render(
        <MemoryRouter>
          <Home />
        </MemoryRouter>
      );
    });
  
    expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);
  
    await act(async () => {
      jest.advanceTimersByTime(499);
    });
  
    expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);
  
    await act(async () => {
      jest.advanceTimersByTime(1);
    });
  
    expect(screen.queryByTestId('loading-indicator')).toBeNull();
    expect(screen.getByText('ユーザー数：10')).toBeInTheDocument();
    expect(screen.getByText('マイクロポスト数：5')).toBeInTheDocument();
  
    jest.useRealTimers();
  });

  test('エラー時はエラーが表示される', async () => {
    jest.useFakeTimers();
  
    axios.get.mockRejectedValue(new Error('Network Error'));
  
    await act(async () => {
      render(
        <MemoryRouter>
          <Home />
        </MemoryRouter>
      );
    });
  
    expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);
  
    await act(async () => {
      jest.advanceTimersByTime(499);
    });
  
    expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);
  
    await act(async () => {
      jest.advanceTimersByTime(1);
    });
  
    expect(screen.queryByTestId('loading-indicator')).toBeNull();
    const userCountElement = screen.getByText('ユーザー数：').closest('p');
    expect(within(userCountElement).getByText('Error: error')).toBeInTheDocument();
    const micropostCountElement = screen.getByText('マイクロポスト数：').closest('p');
    expect(within(micropostCountElement).getByText('Error: error')).toBeInTheDocument();
    jest.useRealTimers();
  });
});
