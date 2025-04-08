✅ Endpoint
GET /api/word-groups/:groupId/words?page=1&limit=10

📥 Request Parameters
| Name     | Location  | Type   | Required | Notes                 |
|----------|-----------|--------|----------|------------------------|
| groupId  | URL Path  | string | ✔️       | ID of the word group   |
| page     | Query     | int    | ❌       | Default is 1           |
| limit    | Query     | int    | ❌       | Default is 10          |

📤 Response

```json
{
  "data": [
    {
      "id": "660e5fb41de3e1a44f9cce91",
      "groupId": "660e5f8b1de3e1a44f9cce90",
      "userId": "660e5f4d1de3e1a44f9cce8f",
      "word": "resilient",
      "language": "English",
      "pronunciations": [
        {
          "ipa": "/rɪˈzɪliənt/",
          "audioUrl": "https://example.com/audio/resilient.mp3"
        }
      ],
      "meanings": [
        {
          "part_of_speech": "adjective",
          "definition": "able to withstand or recover quickly from difficult conditions",
          "examples": [
            "She is very resilient to stress and pressure."
          ]
        }
      ],
      "synonyms": ["tough", "strong"],
      "antonyms": ["fragile"],
      "createdAt": 1712576012
    }
  ],
  "pagination": {
    "page": 1,
    "limit": 10,
    "total": 42
  }
}
```