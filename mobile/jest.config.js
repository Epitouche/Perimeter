module.exports = {
  preset: 'react-native',
  testMatch: ['**/?(*.)+(test|spec).tsx'],
  moduleFileExtensions: ['tsx', 'ts', 'js', 'jsx'],
  testPathIgnorePatterns: ['/node_modules/', '/dist/'],
  transformIgnorePatterns: [
    'node_modules/(?!react-native|react-native-app-auth|@react-native|@react-navigation)',
  ],
};
