import { configureStore } from "@reduxjs/toolkit";
import mentorReducer from "./people/mentorSlice";
import studentReducer from "./people/studentSlice";

const store = configureStore({
  reducer: {
    mentors: mentorReducer,
    students: studentReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
