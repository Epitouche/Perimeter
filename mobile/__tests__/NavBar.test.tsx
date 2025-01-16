/**
 * @file NavBar.test.tsx
 * @description This file contains unit tests for the BottomNavBar component using @testing-library/react-native.
 */

import React from 'react';
import { render, fireEvent } from '@testing-library/react-native';
import { NavigationContainer } from '@react-navigation/native';
import BottomNavBar from '../src/screens/NavBar';

describe('BottomNavBar', () => {
  /**
   * Mock function for navigation.
   */
  const navigate = jest.fn(); // Mock the navigate function

  /**
   * Helper function to render the BottomNavBar component within a NavigationContainer.
   * @returns {RenderResult} The result of rendering the component.
   */
  const renderComponent = () =>
    render(
      <NavigationContainer>
        <BottomNavBar navigation={{ navigate } as any} />
      </NavigationContainer>,
    );

  it('should render the navigation bar', () => {
    const { getByLabelText } = renderComponent();

    // Check if all navigation buttons are rendered
    expect(getByLabelText('Home navigation button')).toBeTruthy();
    expect(getByLabelText('Workflow navigation button')).toBeTruthy();
    expect(getByLabelText('Service navigation button')).toBeTruthy();
    expect(getByLabelText('Settings navigation button')).toBeTruthy();
  });

  it('should navigate to AreaView when Home button is pressed', () => {
    const { getByLabelText } = renderComponent();

    // Simulate pressing the Home button
    fireEvent.press(getByLabelText('Home navigation button'));
    // Check if navigate function is called with 'AreaView'
    expect(navigate).toHaveBeenCalledWith('AreaView');
  });

  it('should navigate to WorkflowScreen when Workflow button is pressed', () => {
    const { getByLabelText } = renderComponent();

    // Simulate pressing the Workflow button
    fireEvent.press(getByLabelText('Workflow navigation button'));
    // Check if navigate function is called with 'WorkflowScreen'
    expect(navigate).toHaveBeenCalledWith('WorkflowScreen');
  });

  it('should navigate to ServicesScreen when Service button is pressed', () => {
    const { getByLabelText } = renderComponent();

    // Simulate pressing the Service button
    fireEvent.press(getByLabelText('Service navigation button'));
    // Check if navigate function is called with 'ServicesScreen'
    expect(navigate).toHaveBeenCalledWith('ServicesScreen');
  });

  it('should navigate to SettingsScreen when Settings button is pressed', () => {
    const { getByLabelText } = renderComponent();

    // Simulate pressing the Settings button
    fireEvent.press(getByLabelText('Settings navigation button'));
    // Check if navigate function is called with 'SettingsScreen'
    expect(navigate).toHaveBeenCalledWith('SettingsScreen');
  });
});
