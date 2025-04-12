-- tasks
create table if not exists tasks (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references auth.users(id) on delete cascade,
  title text not null,
  description text,
  is_completed boolean default false,
  due_date timestamp,
  created_at timestamp default now()
);

-- pomodoro_sessions
create table if not exists pomodoro_sessions (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references auth.users(id) on delete cascade,
  task_id uuid references tasks(id),
  started_at timestamp not null,
  ended_at timestamp,
  is_interrupted boolean default false,
  break_duration integer default 5
);

-- interrupt_logs
create table if not exists interrupt_logs (
  id uuid primary key default gen_random_uuid(),
  session_id uuid references pomodoro_sessions(id) on delete cascade,
  reason text,
  timestamp timestamp default now()
);

-- optimization_profiles
create table if not exists optimization_profiles (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references auth.users(id) on delete cascade,
  recommended_work_duration integer,
  recommended_break_duration integer,
  updated_at timestamp default now()
);

-- notification_settings
create table if not exists notification_settings (
  id uuid primary key default gen_random_uuid(),
  user_id uuid references auth.users(id) on delete cascade,
  enable_push boolean default false,
  enable_email boolean default true,
  quiet_hours_start time,
  quiet_hours_end time
);
