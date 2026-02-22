export default function Home() {
  return (
    <section className="w-full max-w-md">
      <div className="rounded-2xl border border-slate-200 bg-white p-8 shadow-xl">
        <div className="mb-6 text-center">
          <p className="text-sm font-semibold tracking-wide text-slate-500">GO AUTH</p>
          <h1 className="mt-2 text-3xl font-bold text-slate-900">Welcome back</h1>
          <p className="mt-2 text-sm text-slate-600">Sign in to continue to your dashboard</p>
        </div>

        <div className="space-y-3">
          <a
            href="http://127.0.0.1:9000/auth/github/login"
            className="flex w-full items-center justify-center rounded-lg bg-slate-900 px-4 py-3 text-sm font-semibold text-white transition hover:bg-slate-800"
          >
            Continue with GitHub
          </a>

          <a
            href="http://127.0.0.1:9000/auth/google/login"
            className="flex w-full items-center justify-center rounded-lg border border-slate-300 bg-white px-4 py-3 text-sm font-semibold text-slate-700 transition hover:bg-slate-50"
          >
            Continue with Google
          </a>
        </div>

        <p className="mt-6 text-center text-xs text-slate-500">
          By continuing, you agree to secure sign-in using OAuth and JWT.
        </p>
      </div>
    </section>
  );
}
