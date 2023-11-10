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

      let userCountElement = screen.getByTestId('user-count');
      let micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(499);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(1);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).queryByTestId('loading-indicator')).toBeNull();
      expect(within(micropostCountElement).queryByTestId('loading-indicator')).toBeNull();
      expect(within(userCountElement).getByText(/10/)).toBeInTheDocument();
      expect(within(micropostCountElement).getByText(/5/)).toBeInTheDocument();
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

      let userCountElement = screen.getByTestId('user-count');
      let micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(499);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(1);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).queryByTestId('loading-indicator')).toBeNull();
      expect(within(micropostCountElement).queryByTestId('loading-indicator')).toBeNull();
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

      let userCountElement = screen.getByTestId('user-count');
      let micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(1999);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).getByTestId('loading-indicator')).toBeInTheDocument();
      expect(within(micropostCountElement).getByTestId('loading-indicator')).toBeInTheDocument();

      await act(async () => {
        jest.advanceTimersByTime(1);
      });

      userCountElement = screen.getByTestId('user-count');
      micropostCountElement = screen.getByTestId('micropost-count');
      expect(within(userCountElement).queryByTestId('loading-indicator')).toBeNull();
      expect(within(micropostCountElement).queryByTestId('loading-indicator')).toBeNull();
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
