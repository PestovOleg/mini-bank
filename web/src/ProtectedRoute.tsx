import { observer } from "mobx-react-lite";
import store from "./store/store";
import { Navigate, useLocation } from "react-router-dom";

const ProtectedRoute = ({ children }: { children: JSX.Element }) => {
    
    let location = useLocation();
    let isAuth=store.userStore.isAuth
    console.log(isAuth)
    if (!isAuth) {
  
      return <Navigate to="/login" state={{ from: location }} replace />;
    }
  
    return children;
  }
const ProtectedRouteObserver = observer(ProtectedRoute);

export default ProtectedRouteObserver;