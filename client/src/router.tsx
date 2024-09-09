import { createBrowserRouter } from "react-router-dom";
import { Suspense } from "react";
import { Spin } from "antd";
import Auth from "./components/auth/Auth";
import NotFound from "./components/error/NotFound";
import Dashboard from "./components/dashboard/Dashboard";
import UserCourses from "./components/courses/list/UserCourses";
import CourseOverview from "./components/courses/CourseOverview";

import CreateCourseComponent from "./components/courses/create/CreateCourseComponent";

export const router = createBrowserRouter([
  {
    path: "/login",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <Auth />
      </Suspense>
    ),
  },
  {
    path: "",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <Dashboard />
      </Suspense>
    ),
  },
  {
    path: "/courses",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <UserCourses />
      </Suspense>
    ),
  },
  {
    path: "/courses/create",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <CreateCourseComponent />
      </Suspense>
    ),
  },
  {
    path: "/courses/:courseId",
    element: (
      <Suspense fallback={<Spin size="large" />}>
        <CourseOverview />
      </Suspense>
    ),
  },
  {
    path: "*",
    element: <NotFound />,
  },
]);
