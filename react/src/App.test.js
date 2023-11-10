import React from 'react';
import { render, screen, act } from '@testing-library/react';
import { MemoryRouter } from 'react-router-dom';
import Home from './components/Home';
import Users from './components/Users';
import Microposts from './components/Microposts';

jest.mock('axios');

describe('App Router', () => {
  test('Home component is rendered at /', async () => {
    window.history.pushState({}, '', '/');
    await act(async () => {
      render(
        <MemoryRouter>
          <Home />
        </MemoryRouter>,
      );
    });
    const homeElement = await screen.findByTestId('home');
    expect(homeElement).toBeInTheDocument();
  });

  test('Users component is rendered at /users', async () => {
    window.history.pushState({}, '', '/users');
    await act(async () => {
      render(
        <MemoryRouter>
          <Users />
        </MemoryRouter>,
      );
    });
    const usersElement = await screen.findByTestId('users');
    expect(usersElement).toBeInTheDocument();
  });

  test('Microposts component is rendered at /microposts', async () => {
    window.history.pushState({}, '', '/microposts');
    await act(async () => {
      render(
        <MemoryRouter>
          <Microposts />
        </MemoryRouter>,
      );
    });
    const micropostsElement = await screen.findByTestId('microposts');
    expect(micropostsElement).toBeInTheDocument();
  });
});
