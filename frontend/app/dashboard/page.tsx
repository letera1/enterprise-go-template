export default function DashboardPage() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center bg-gray-100 p-6">
      <div className="w-full max-w-4xl rounded-xl bg-white p-8 shadow-lg">
        <h1 className="mb-6 text-3xl font-bold text-gray-800">Dashboard</h1>
        <p className="mb-4 text-gray-600">
          Welcome to your protected dashboard! This data is secured by the interactions
          between Next.js and your Go backend (PostgreSQL).
        </p>
        <div className="rounded-lg bg-blue-50 p-4 border border-blue-200">
          <p className="text-sm text-blue-800">
            <strong>Note:</strong> To see real data here, your Go backend needs to
            expose a protected API endpoint that this page can fetch from using the
            session token / cookie.
          </p>
        </div>
        <div className="mt-8">
           <a href="/" className="text-indigo-600 hover:underline">
             &larr; Back to Home
           </a>
        </div>
      </div>
    </div>
  );
}
