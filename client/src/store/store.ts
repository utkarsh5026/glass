import { configureStore } from "@reduxjs/toolkit";
import mentorReducer from "./people/mentorSlice";
import studentReducer from "./people/studentSlice";
import assignmentReducer from "./activity/assignmentSlice";

const store = configureStore({
  reducer: {
    mentors: mentorReducer,
    students: studentReducer,
    assignments: assignmentReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
