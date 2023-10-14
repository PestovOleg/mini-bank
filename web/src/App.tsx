import React, { useEffect } from 'react';
import './App.css';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import ProtectedRouteObserver from './ProtectedRoute';
import MainPageObserver from './MainPage';
import LoginObserver from './components/Login';
import store from './store/store';

function App() {  
  useEffect(() => {
    store.toggleStore.getToggles()
  }, []);
  return (
    <Router>
      <Routes>
        <Route path='/' element={
          <ProtectedRouteObserver>
            <MainPageObserver />
          </ProtectedRouteObserver>
        } />
        <Route path="/login" element={<LoginObserver />} />
      </Routes>
    </Router>
  );
}

export default App;
