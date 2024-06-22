import React from 'react';
import { Redirect, Route, RouteProps } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { RootState } from '../../redux/store';

interface ProtectedRouteProps extends RouteProps {
  component: React.ComponentType<any>;
}

const ProtectedRoute: React.FC<ProtectedRouteProps> = ({ component: Component, ...rest }) => {
  const token = localStorage.getItem('token');
  const user = useSelector((state: RootState) => state.user);

  if (!token) {
    return <Redirect to="/login" />;
  }

  const decodedToken = JSON.parse(atob(token.split('.')[1]));
  const { role, idnumber } = decodedToken;

  return (
    <Route
      {...rest}
      render={props =>
        (rest.path.startsWith("/admin") && role !== "Admin") ||
        (rest.path.startsWith("/user-dashboard") && props.match.params.idnumber !== idnumber) ? (
          <Redirect to="/unauthorized" />
        ) : (
          <Component {...props} />
        )
      }
    />
  );
};

export default ProtectedRoute;
