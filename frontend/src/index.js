import React from 'react';
import ReactDOM from 'react-dom';
import { Router, Route, browserHistory } from 'react-router';
import App from './components/App';
import CarsCatalog from './components/CarsCatalog';
import Agents from './components/Agents';
import Agencies from './components/Agencies';
import Status from './components/Status';

ReactDOM.render(
  <Router history={browserHistory}>
    <Route path="/" component={App} />
    <Route path="/cars" component={CarsCatalog} />
    <Route path="/agents" component={Agents} />
    <Route path="/agencies" component={Agencies} />
    <Route path="/status" component={Status} />
  </Router>,
  document.getElementById('root')
);
