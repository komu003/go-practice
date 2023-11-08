import React from 'react';
import { render, screen, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import axios from 'axios';
import Home from './Home';
import { MemoryRouter } from 'react-router-dom';

jest.mock('axios');

describe('Home コンポーネント', () => {
  test('初期状態でローディングインジケーターが表示される', async () => {
    axios.get.mockImplementation(() => new Promise(() => {}));
    render(
      <MemoryRouter>
        <Home />
      </MemoryRouter>
    );
    const loadingIndicators = screen.getAllByTestId('loading-indicator');
    expect(loadingIndicators.length).toBeGreaterThan(0);
  });
});
