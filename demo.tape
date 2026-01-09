# RagTune Demo - "See What Your RAG Retrieves"
# Run with: vhs demo.tape

Output assets/demo.gif

Set Shell "bash"
Set FontSize 14
Set Width 1000
Set Height 600
Set Theme "Catppuccin Mocha"
Set Padding 10

# Setup (hidden)
Hide
Type "cd /Users/user/Projects/ragtune"
Enter
Sleep 500ms
Show

# Title
Type "# RagTune — See what your RAG retrieves"
Enter
Sleep 1500ms

# Scene 1: explain - Debug a query
Type "./ragtune explain 'How do I create an API key?' --collection demo-1024 --embedder ollama"
Enter
Sleep 6s

Type ""
Enter
Type "# Score 0.77 — looks reasonable"
Enter
Sleep 3s

# Scene 2: audit - Batch health check
Type ""
Enter
Type "# Let's check overall health..."
Enter
Sleep 1500ms

Type "./ragtune audit --collection demo-128 --queries ./data/queries.json --embedder ollama"
Enter
Sleep 6s

Type ""
Enter
Type "# Recall 0.68 — below threshold. Worth investigating."
Enter
Sleep 3s

# Scene 3: compare - Measure the difference
Type ""
Enter
Type "# How much does chunk size matter here?"
Enter
Sleep 1500ms

Type "./ragtune compare --collections demo-128,demo-1024 --queries ./data/queries.json --embedder ollama"
Enter
Sleep 7s

Type ""
Enter
Type "# 18% difference. Now you have data, not guesses."
Enter
Sleep 4s
