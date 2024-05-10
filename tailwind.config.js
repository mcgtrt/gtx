/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./**/*.html", 
    "./**/*.templ", 
    "./**/*.go",
  ],
  theme: {
    extend: {},
    colors: {
      primary: '#2E90FA',
      secondary: '#1C75BC',
      accent: '#F2994A',
      inactive: '#F2F2F2',
      success: '#27AE60',
      warning: '#E74C3C',
      text3: '#333333',
      text6: '#666666',
      text9: '#999999',
      textc: "#CCCCCC",
      white: '#FFFFFF',
    },
    textColor: {
      primary: '#2E90FA',
      secondary: '#1C75BC',
      accent: '#F2994A',
      inactive: '#F2F2F2',
      success: '#27AE60',
      warning: '#E74C3C',
      text3: '#333333',
      text6: '#666666',
      text9: '#999999',
      textc: "#CCCCCC",
      white: '#FFFFFF',
    },
  },
  plugins: [],
}

