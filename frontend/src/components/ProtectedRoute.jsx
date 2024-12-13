import { useSelector } from "react-redux";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
    const { accessToken } = useSelector((state) => state.accessToken);

    return accessToken?.length > 0 ? <>{children}</> : <Navigate to="/" />;
};

export default ProtectedRoute