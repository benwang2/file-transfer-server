existing = set()
bad = set()

with open("words.txt","r") as f:
    existing.update([l.strip() for l in f.readlines()])
    osize = len(existing)

with open("bad.txt","r") as f:
    bad.update([l.strip() for l in f.readlines()])

with open("5.txt","r") as f:
    c = ""
    while c != "q":
        word = f.readline().strip()
        if (word in existing) or (word in bad): continue
        try:
            c = input(word+": ")
            if c == "y":
                existing.add(word)
            else:
                bad.add(word.strip())
        except KeyboardInterrupt:
            break

with open("words.txt","w") as f:
    print(f"Added {len(existing)-osize} words.")
    s = sorted(existing, key=lambda x:(len(x), x))
    f.write("\n".join(s))

with open("bad.txt","w") as f:
    f.write("\n".join(bad))