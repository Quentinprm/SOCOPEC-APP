import React from 'react';
import '../style/topNavBarStyle.css';
import { Navbar, Nav, NavItem } from 'react-bootstrap/lib';

function TopNavbar() {
  return (
    <div className="topnav">
      <a className="active" href="/">
        Socopec.App
      </a>
      <a href="/cars">Véhicules</a>
      <a href="/agents">Agents</a>
      <a href="/agencies">Agences</a>
      <a href="/status">Status</a>
    </div>
  );
}

export default TopNavbar;
