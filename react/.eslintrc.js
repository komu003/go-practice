module.exports = {
  env: {
    browser: true,
    es2021: true,
  },
  extends: 'airbnb',
  overrides: [
    {
      env: {
        node: true,
      },
      files: [
        '.eslintrc.{js,cjs}',
      ],
      parserOptions: {
        sourceType: 'script',
      },
    },
  ],
  parserOptions: {
    ecmaVersion: 'latest',
    sourceType: 'module',
  },
  rules: {
    'react/function-component-definition': [
      2,
      {
        namedComponents: 'arrow-function',
        unnamedComponents: 'arrow-function',
      },
    ],
    'react/jsx-filename-extension': [1, { extensions: ['.js', '.jsx'] }],
  },
};
