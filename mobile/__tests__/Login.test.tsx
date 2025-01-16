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
        const { getByPlaceholderText, getByText } = render(<LoginScreen navigation={MockNavigation} route={MockRoute} />);
        expect(getByPlaceholderText('Email')).toBeTruthy();
        expect(getByPlaceholderText('Password')).toBeTruthy();
        expect(getByText('Login')).toBeTruthy();
    });
    
    it('should show error message when email is empty', () => {
        const { getByText, getByPlaceholderText } = render(<LoginScreen navigation={MockNavigation} route={MockRoute} />);
        fireEvent.changeText(getByPlaceholderText('Email'), '');
        fireEvent.changeText(getByPlaceholderText('Password'), 'password');
        fireEvent.press(getByText('Login'));
        expect(getByText('Email is required')).toBeTruthy();
    });

    it('should show error message when password is empty', () => {
        const { getByText, getByPlaceholderText } = render(<LoginScreen navigation={MockNavigation} route={MockRoute} />);
        fireEvent.changeText(getByPlaceholderText('Email'), 'test@example.com');
        fireEvent.changeText(getByPlaceholderText('Password'), '');
        fireEvent.press(getByText('Login'));
        expect(getByText('Password is required')).toBeTruthy();
    });
        
    it('should call login function with correct credentials', () => {
        const mockLogin = jest.fn();
        const { getByText, getByPlaceholderText } = render(<LoginScreen navigation={MockNavigation} route={MockRoute} />);
        fireEvent.changeText(getByPlaceholderText('Email'), 'test@example.com');
        fireEvent.changeText(getByPlaceholderText('Password'), 'password');
        fireEvent.press(getByText('Login'));
        expect(mockLogin).toHaveBeenCalledWith('test@example.com', 'password');
    });
});