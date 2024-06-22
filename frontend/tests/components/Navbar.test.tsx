import { it, expect, describe } from 'vitest'
import { render , screen , fireEvent, waitFor} from "@testing-library/react";
import Navbar from '../../src/components/Navbar/Navbar'
import { Provider } from 'react-redux';
import { MemoryRouter } from 'react-router-dom';
import store from '../../src/redux/store';


describe('navbar', () => {
    it('display username at the navbar if there is username in redux store', async () => {
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
                <Navbar />
              </MemoryRouter>
            </Provider>
          );

        const usernameElement = screen.getByText("getter")
        expect(usernameElement).toBeTruthy(); //// เหมือนๆกัน, use .toBeDefined() or .toBeVisible()
    });
    it('display logout button if there is username in redux store', async () => {
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
              <Navbar />
            </MemoryRouter>
          </Provider>
        );

      const button = screen.getByRole("button",{name : /LogOut/i})
      expect(button).toBeTruthy(); //// เหมือนๆกัน, use .toBeDefined() or .toBeVisible()
  })
  it('display login button if there is no username in redux store', async () => {
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
            <Navbar />
          </MemoryRouter>
        </Provider>
      );

    const button = screen.getByRole("button",{name : /Log In/i})
    expect(button).toBeTruthy(); //// เหมือนๆกัน, use .toBeDefined() or .toBeVisible()
  })
  it('display signup button if there is no username in redux store', async () => {
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
            <Navbar />
          </MemoryRouter>
        </Provider>
      );

    const button = screen.getByRole("button",{name : /Sign Up/i})
    expect(button).toBeTruthy(); //// เหมือนๆกัน, use .toBeDefined() or .toBeVisible()
  })
})