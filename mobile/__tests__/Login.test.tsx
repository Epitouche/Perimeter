import React from 'react';
import { render, fireEvent } from '@testing-library/react-native';
import LoginScreen from '../src/screens/LoginScreen';
import { AppProvider } from '../src/context/AppContext';

const MockNavigation = {
  navigate: jest.fn(),
  dispatch: jest.fn(),
  goBack: jest.fn(),
  reset: jest.fn(),
  setParams: jest.fn(),
  setOptions: jest.fn(),
  isFocused: jest.fn(),
  canGoBack: jest.fn(),
  getParent: jest.fn(),
  getState: jest.fn(),
  navigateDeprecated: jest.fn(),
  preload: jest.fn(),
  getId: jest.fn(),
  setStateForNextRouteNamesChange: jest.fn(),
  addListener: jest.fn(),
  removeListener: jest.fn(),
  replace: jest.fn(),
  push: jest.fn(),
  pop: jest.fn(),
  popToTop: jest.fn(),
  popTo: jest.fn(),
};

jest.mock('../src/context/AppContext', () => {
  const originalModule = jest.requireActual('../src/context/AppContext');
  return {
    ...originalModule,
    AppProvider: ({ children }: { children: React.ReactNode }) => (
      <originalModule.AppContext.Provider value={{ ipAddress: 'localhost'}}>
        {children}
      </originalModule.AppContext.Provider>
    ),
  };
});

const MockRoute = { key: 'login', name: 'Login' as const, params: undefined };

describe('LoginScreen', () => {
  it('should render correctly', () => {
    const { getByPlaceholderText } = render(
      <AppProvider>
        <LoginScreen navigation={MockNavigation} route={MockRoute} />
      </AppProvider>
    );
    expect(getByPlaceholderText('Enter username')).toBeTruthy();
    expect(getByPlaceholderText('Enter password')).toBeTruthy();
  });

  it('should show error message when email is empty', () => {
    const { getByText, getByPlaceholderText, getByTestId } = render(
      <AppProvider>
        <LoginScreen navigation={MockNavigation} route={MockRoute} />
      </AppProvider>
    );
    fireEvent.changeText(getByPlaceholderText('Enter username'), '');
    fireEvent.changeText(getByPlaceholderText('Enter password'), 'password');
    fireEvent.press(getByTestId('login-button'));
    expect(getByText('Username is required')).toBeTruthy();
  });

  it('should show error message when password is empty', () => {
    const { getByText, getByPlaceholderText, getByTestId } = render(
      <AppProvider>
        <LoginScreen navigation={MockNavigation} route={MockRoute} />
      </AppProvider>
    );
    fireEvent.changeText(getByPlaceholderText('Enter username'), 'username');
    fireEvent.changeText(getByPlaceholderText('Enter password'), '');
    fireEvent.press(getByTestId('login-button'));
    expect(getByText('Password is required')).toBeTruthy();
  });
});
