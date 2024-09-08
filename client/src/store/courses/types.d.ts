export interface Course {
  id: string;
  title: string;
  creator: string;
  subject: string;
  description: string;
  startDate: string;
  endDate: string;
  category: string;
}

export interface CourseBasic {
  id: string;
  title: string;
  creator: string;
  subject: string;
  description: string;
}
