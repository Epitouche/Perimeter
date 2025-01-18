import React from 'react';
import { render, fireEvent } from '@testing-library/react-native';
import LoginScreen from '../src/screens/LoginScreen';

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
const MockRoute = { key: 'login', name: 'Login' as const, params: undefined };

describe('LoginScreen', () => {
  it('should render correctly', () => {
    const { getByPlaceholderText } = render(
      <LoginScreen navigation={MockNavigation} route={MockRoute} />,
    );
    expect(getByPlaceholderText('Enter username')).toBeTruthy();
    expect(getByPlaceholderText('Enter password')).toBeTruthy();
  });

  it('should show error message when email is empty', () => {
    const { getByText, getByPlaceholderText } = render(
      <LoginScreen navigation={MockNavigation} route={MockRoute} />,
    );
    fireEvent.changeText(getByPlaceholderText('Enter username'), '');
    fireEvent.changeText(getByPlaceholderText('Enter password'), 'password');
    fireEvent.press(getByText('Log in'));
    expect(getByText('Username is required')).toBeTruthy();
  });

  it('should show error message when password is empty', () => {
    const { getByText, getByPlaceholderText } = render(
      <LoginScreen navigation={MockNavigation} route={MockRoute} />,
    );
    fireEvent.changeText(getByPlaceholderText('Enter username'), 'username');
    fireEvent.changeText(getByPlaceholderText('Enter password'), '');
    fireEvent.press(getByText('Log in'));
    expect(getByText('Password is required')).toBeTruthy();
  });

  it('should call login function with correct credentials', () => {
    const mockLogin = jest.fn();
    const { getByText, getByPlaceholderText } = render(
      <LoginScreen navigation={MockNavigation} route={MockRoute} />,
    );
    fireEvent.changeText(getByPlaceholderText('Enter username'), 'username');
    fireEvent.changeText(getByPlaceholderText('Enter password'), 'password');
    fireEvent.press(getByText('Log in'));
    expect(mockLogin).toHaveBeenCalledWith('test@example.com', 'password');
  });
});
