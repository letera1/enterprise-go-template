'use client';

import { useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';

type MeResponse = {
  user_id: string;
  email: string;
  message: string;
};

export default function Dashboard() {
  const [user, setUser] = useState<MeResponse | null>(null);
  const [loading, setLoading] = useState(true);
  const router = useRouter();

  useEffect(() => {
    fetch('http://127.0.0.1:9000/api/me', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Important to send cookies
    })
      .then((res) => {
        if (res.status === 401) {
          router.push('/');
          return null;
        }
        return res.json();
      })
      .then((data) => {
        if (data) {
          setUser(data);
        }
        setLoading(false);
      })
      .catch((err) => {
        console.error(err);
        setLoading(false);
      });
  }, [router]);

  const handleLogout = () => {
    fetch('http://127.0.0.1:9000/auth/logout', { 
        method: 'POST',
        credentials: 'include'
    }).then(() => {
        router.push('/');
    });
  };

  if (loading) {
    return (
      <section className="w-full max-w-3xl">
        <div className="rounded-2xl border border-slate-200 bg-white p-8 shadow-xl">
          <p className="text-sm text-slate-500">Loading your dashboard...</p>
        </div>
      </section>
    );
  }

  return (
    <section className="w-full max-w-3xl">
      <div className="rounded-2xl border border-slate-200 bg-white p-8 shadow-xl">
        <div className="mb-8 flex items-center justify-between">
          <div>
            <p className="text-sm font-semibold tracking-wide text-slate-500">GO AUTH</p>
            <h1 className="mt-1 text-3xl font-bold text-slate-900">Dashboard</h1>
          </div>
          <button
            onClick={handleLogout}
            className="rounded-lg bg-slate-900 px-4 py-2 text-sm font-semibold text-white transition hover:bg-slate-800"
          >
            Logout
          </button>
        </div>

        <div className="mb-6 rounded-lg border border-emerald-200 bg-emerald-50 p-4">
          <h2 className="text-lg font-semibold text-emerald-800">Welcome back</h2>
          <p className="mt-1 text-sm text-emerald-700">You are signed in successfully.</p>
        </div>

        <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
          <div className="rounded-lg border border-slate-200 bg-slate-50 p-4">
            <p className="text-xs font-semibold uppercase tracking-wide text-slate-500">User ID</p>
            <p className="mt-2 break-all font-mono text-sm text-slate-800">{user?.user_id}</p>
          </div>
          <div className="rounded-lg border border-slate-200 bg-slate-50 p-4">
            <p className="text-xs font-semibold uppercase tracking-wide text-slate-500">Email</p>
            <p className="mt-2 break-all font-mono text-sm text-slate-800">{user?.email}</p>
          </div>
        </div>

        <p className="mt-6 text-xs text-slate-500">
          Session is secured with OAuth and JWT in an HttpOnly cookie.
        </p>
      </div>
    </section>
  );
}
