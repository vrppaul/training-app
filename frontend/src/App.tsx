import { FC } from 'react';
import {
  BrowserRouter as Router,
  Link, Route, Switch,
} from 'react-router-dom';
import ExercisePage from './components/exercises';

const Home: FC = () => (
  <button type="button">
    <Link to="/exercises">My exercises</Link>
  </button>
);

const App: FC = () => (
  <Router>
    <Switch>
      <Route exact path="/" component={Home} />
      <Route path="/exercises" component={ExercisePage} />
    </Switch>
  </Router>
);

export default App;
