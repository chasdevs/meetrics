-- Time spent in meetings
SET @minsInWorkDay = 7*60; # 7 hours of "work time"
SELECT
  umm.date,
  AVG(umm.mins2_plus/@minsInWorkDay*100) AS "% time in meetings",
  AVG(umm.mins1/@minsInWorkDay*100) AS "% time in 1:1s",
  AVG((CAST(@minsInWorkDay AS SIGNED) - CAST(umm.mins1 + umm.mins2_plus AS SIGNED))/@minsInWorkDay*100) AS "% crank time"
FROM user_meeting_mins umm
JOIN users u
  ON umm.user_id = u.id
GROUP BY YEARWEEK(umm.date);


-- Avg meeting time by department
SET @minsInWorkDay = 7*60; # 7 hours of "work time"
SELECT
  u.department,
  AVG(umm.mins2_plus/@minsInWorkDay*100) AS "% time in meetings",
  AVG(umm.mins1/@minsInWorkDay*100) AS "% time in 1:1s",
  AVG((CAST(@minsInWorkDay AS SIGNED) - CAST(umm.mins1 + umm.mins2_plus AS SIGNED))/@minsInWorkDay*100) AS "% crank time"
FROM user_meeting_mins umm
JOIN users u
  ON umm.user_id = u.id
WHERE u.department NOT IN ("Enterprise")
GROUP BY u.department;

-- Biggest recurring meetings
SELECT
  m.name,
  m.attendees,
  m.mins,
  m.frequency_per_month * m.mins * m.attendees / 60 AS hrs_per_month
FROM meetings m
WHERE m.name NOT LIKE "%football%" AND m.name NOT LIKE "%softball%"
ORDER BY hrs_per_month DESC
LIMIT 5;