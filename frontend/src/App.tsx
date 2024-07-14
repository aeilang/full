import "./App.css";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { DashBoard, Home } from "./page";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
  },
  {
    path: "/dashboard",
    element: <DashBoard />,
  },
  {
    path: "*",
    element: <h1>Not Found</h1>,
  },
]);

function App() {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
}

export default App;
