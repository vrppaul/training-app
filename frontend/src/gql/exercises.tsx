import {
  gql,
} from '@apollo/client';

interface Exercise {
  _id: string;
  name: string;
  description: string;
}

export interface ExercisesData {
  exercises: Exercise[];
}

export const GET_EXERCISES = gql`
  query GetExercises {
    exercises {
      _id
      name
      description
    }
  }
`;

export const CREATE_EXERCISE = gql`
  mutation CreateExercise($name: String!, $description: String!) {
    createExercise(input: {name: $name, description: $description}) {
      _id
      name
      description
    }
  }
`;
