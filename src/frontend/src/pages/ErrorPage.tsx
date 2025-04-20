import { useRouteError } from 'react-router-dom';

export default function ErrorPage() {
  const error = useRouteError();
  
  return (
    <div className="flex flex-col items-center justify-center min-h-screen">
      <h1 className="text-2xl font-bold mb-4">Oops! Something went wrong</h1>
      <p className="text-gray-600">
        {error instanceof Error ? error.message : 'Please try again later'}
      </p>
    </div>
  );
}