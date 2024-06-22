// src/components/Greet/Greet.test.tsx
import { it, expect, describe } from 'vitest';
import { render, screen } from '@testing-library/react';
import '@testing-library/jest-dom/vitest';
import Home from '../../src/pages/Home/Home'
import { Provider } from 'react-redux';
import { MemoryRouter } from 'react-router-dom';
import store from '../../src/redux/store';

describe('home', () => {
  it('the button should be changed to "Go Admin Dashboard" if current user is admin', () => {
    store.dispatch({
      type: 'user/login',
      payload: {
        idnumber: '123',
        email: 'test@example.com',
        username: 'getter',
        role: 'Admin',
      },
    });
    render(
      <Provider store={store}>
        <MemoryRouter>
          <Home />
        </MemoryRouter>
      </Provider>
    );

    const button = screen.getByRole('button', {name:/Go Admin Dashboard/i});
    expect(button).toBeInTheDocument();
  });
  it('the button should be changed to "Go User Dashboard" if current user is user', () => {
    store.dispatch({
      type: 'user/login',
      payload: {
        idnumber: '123',
        email: 'test@example.com',
        username: 'getter',
        role: 'User',
      },
    });
    render(
      <Provider store={store}>
        <MemoryRouter>
          <Home />
        </MemoryRouter>
      </Provider>
    );

    const button = screen.getByRole('button', {name:/Go User Dashboard/i});
    expect(button).toBeInTheDocument();
  });
  it('the button should be changed to "Go User Dashboard" if current user is user', () => {
    store.dispatch({
      type: 'user/login',
      payload: {
        idnumber: '',
        email: '',
        username: '',
        role: '',
      },
    });
    render(
      <Provider store={store}>
        <MemoryRouter>
          <Home />
        </MemoryRouter>
      </Provider>
    );

    const button = screen.getByRole('button', {name:/Sign Up/i});
    expect(button).toBeInTheDocument();
  });
});
