import React, {
  Dispatch, FC, SetStateAction, useState,
} from 'react';
import {
  useQuery,
  useMutation,
} from '@apollo/client';
import {
  CREATE_EXERCISE, GET_EXERCISES, ExercisesData, GET_EXERCISE_FRAGMENT,
} from '../gql/exercises';
import Loading from './loading';

interface ExerciseProps {
  name: string;
  description: string;
}

interface NewExerciseProps {
  setDisplayExisting: Dispatch<SetStateAction<boolean>>;
}

const Exercise: FC<ExerciseProps> = ({ name, description }: ExerciseProps) => (
  <>
    <b>{name}</b>
    <p>{description}</p>
  </>
);

const ExistingExercises: FC = () => {
  const { loading, error, data } = useQuery<ExercisesData>(GET_EXERCISES);
  return (
    <>
      <h1>
        Мои упражнения
      </h1>
      {loading ? <Loading /> : data?.exercises.map(
        (exercise) => (
          <Exercise
            key={exercise._id}
            name={exercise.name}
            description={exercise.description}
          />
        ),
      )}
      {/* TODO: add sentry */}
      {/* TODO: add correct error handling with error.graphQLErrors and error.networkError */}
      {error && <p>Error occurred, try again</p>}
    </>
  );
};

const NewExercise: FC<NewExerciseProps> = (
  { setDisplayExisting }: NewExerciseProps,
) => {
  const [name, setName] = useState<string>();
  const [description, setDescription] = useState<string>();
  const [createExercise, { loading, error }] = useMutation(CREATE_EXERCISE, {
    update(cache, { data: { createExercise } }) {
      cache.modify({
        fields: {
          exercises(existingExercises = []) {
            const newExerciseRef = cache.writeFragment({
              data: createExercise,
              fragment: GET_EXERCISE_FRAGMENT,
            });
            return [...existingExercises, newExerciseRef];
          },
        },
      });
    },
  });

  const setValue = (
    e: React.FormEvent<HTMLInputElement>,
    setter: React.Dispatch<React.SetStateAction<string | undefined>>,
  ): void => {
    e.preventDefault();
    setter(e.currentTarget.value);
  };

  const createNewExercise = (e: React.FormEvent<HTMLFormElement>): void => {
    e.preventDefault();
    createExercise(
      { variables: { name, description } },
    ).then(() => {
      setDisplayExisting(true);
    }).catch(
      (catchedError) => { console.log(catchedError); },
    );
  };

  return (
    <>
      <h1>
        Создать новое упражнение
      </h1>
      {loading ? <h2>Создаю новое упражнение</h2> : (
        <form onSubmit={createNewExercise}>
          <input
            name="name"
            type="text"
            placeholder="Имя упражнения"
            onChange={
          (e: React.FormEvent<HTMLInputElement>): void => setValue(e, setName)
        }
          />
          <input
            name="description"
            type="textarea"
            placeholder="Описание упражнения"
            onChange={
          (e: React.FormEvent<HTMLInputElement>): void => setValue(e, setDescription)
        }
          />
          <button type="submit">Создать новое упражнение</button>
        </form>
      )}
      {/* TODO: add sentry */}
      {/* TODO: add correct error handling with error.graphQLErrors and error.networkError */}
      {error && <p>Error occurred, try once ogain</p>}
    </>
  );
};

const ExercisePage: FC = () => {
  const [displayExisting, setDisplayExisting] = useState<boolean>(true);

  return (
    <>
      {displayExisting
        ? <ExistingExercises />
        : <NewExercise setDisplayExisting={setDisplayExisting} />}
      <button type="button" onClick={() => { setDisplayExisting(!displayExisting); }}>
        {displayExisting ? 'Создать новое упражнение' : 'Отмена'}
      </button>
    </>
  );
};

export default ExercisePage;
