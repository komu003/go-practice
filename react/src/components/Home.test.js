import React from 'react';
import { render, screen, within } from '@testing-library/react';
import '@testing-library/jest-dom';
import axios from 'axios';
import { MemoryRouter } from 'react-router-dom';
import { act } from 'react-dom/test-utils';
import Home from './Home';
import { API_BASE_URL } from '../config';

jest.mock('axios');

describe('Home コンポーネント', () => {
  describe('APIの応答に応じた表示のテスト', () => {
    beforeEach(() => {
      jest.useFakeTimers();
    });
  
    afterEach(() => {
      jest.useRealTimers();
    });

    test('ユーザー数とマイクロポスト数が表示される', async () => {
      axios.get.mockImplementation((url) => {
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
          </MemoryRouter>,
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
    });

    test('エラー時はエラーが表示される', async () => {
      axios.get.mockRejectedValue(new Error('Network Error'));

      await act(async () => {
        render(
          <MemoryRouter>
            <Home />
          </MemoryRouter>,
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
      const userCountElement = screen.getByTestId('user-count');
      const micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByText(/Error:error/)).toBeInTheDocument();
      expect(within(micropostCountElement).getByText(/Error:error/)).toBeInTheDocument();
    });

    test('タイムアウト時はタイムアウトが表示される', async () => {
      axios.get.mockImplementation(() => new Promise((_, reject) => {
        setTimeout(() => {
          const error = new Error('Network Error');
          error.code = 'ECONNABORTED';
          reject(error);
        }, 2000);
      }));

      await act(async () => {
        render(
          <MemoryRouter>
            <Home />
          </MemoryRouter>,
        );
      });

      expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);

      await act(async () => {
        jest.advanceTimersByTime(1999);
      });

      expect(screen.getAllByTestId('loading-indicator').length).toBeGreaterThan(0);

      await act(async () => {
        jest.advanceTimersByTime(1);
      });

      expect(screen.queryByTestId('loading-indicator')).toBeNull();
      const userCountElement = screen.getByTestId('user-count');
      const micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByText(/Error:timeout/)).toBeInTheDocument();
      expect(within(micropostCountElement).getByText(/Error:timeout/)).toBeInTheDocument();
    });
  });

  test('ユーザー一覧とマイクロポスト一覧へのリンクが存在する', async () => {
    await act(async () => {
      render(
        <MemoryRouter>
          <Home />
        </MemoryRouter>,
      );
    });

    const usersLink = screen.getByRole('link', { name: 'ユーザー一覧' });
    expect(usersLink).toBeInTheDocument();
    expect(usersLink).toHaveAttribute('href', '/users');

    const micropostsLink = screen.getByRole('link', { name: 'マイクロポスト一覧' });
    expect(micropostsLink).toBeInTheDocument();
    expect(micropostsLink).toHaveAttribute('href', '/microposts');
  });
});
