CREATE TABLE feeds (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  desc  text
) STRICT;

CREATE TABLE sightings (
  id INTEGER PRIMARY KEY,
  feed_id INTEGER NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
  created_at TEXT NOT NULL, --ISO DATETIME
  title TEXT,
  summary TEXT,
  lat REAL NOT NULL,
  long REAL NOT NULL
) STRICT;



  --  <entry>
  --     <title>M 3.2, Mona Passage</title>
  --     <link href="http://example.org/2005/09/09/atom01"/>
  --     <id>urn:uuid:1225c695-cfb8-4ebb-aaaa-80da344efa6a</id>
  --     <updated>2005-08-17T07:02:32Z</updated>
  --     <summary>We just had a big one.</summary>
  --     <georss:point>45.256 -71.92</georss:point>
  --  </entry>