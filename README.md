# Hashunter
Hashunter is a proof-of-concept brute-force password cracking tool designed to hunts down passwords by matching their SHA-256 hashes. 
This project demonstrates the fundamental techniques used in brute-force attacks and highlights the importance of strong password practices. 

## How it works
1. Character set: users can define a customizable character set from which the tool will generate potential passwords. By default, it includes lowercase letters and digits.
2. Multi-Threaded Processing: to enhance performance, Hashunter employs multi-threading (through goroutines), allowing it to test multiple password combinations simultaneously.
3. Recursive brute-force algorithm: the core logic is based on a recursive algorithm that builds potential passwords up to a specificed maximum length and checks each one against the target.
4. Timeout mechanism: users can specify a timeout duration, after which the tool will stop searching (testing scenarios).

## Limitations of brute-force attacks 
While Hashunter demonstrates the brute-force method, it's important to recognize its limitations: 
- Time complexity: brute-force approach can be extremely slow, especially with longer passwords and larger character sets. The time required grows exponentially with its length.
- Inefficiency: brute-force checks every possible combination, making it inefficient compared to other techniques.
- Resource intensive: high computational resources are needed for longer password lengths (impractical)

In contrast to brute-force attacks, dictionary attacks and rainbow tables offer more efficient methods for password cracking. 
Dictionary attacks: these attacks use a pre-defined list of potential passwords to guess passwords. 
Rainbow tables: pre-computed tables of hash values for common passwords. By comparing hashes from these tables to the target hash, attackers can quickly identify the original password without having to compute each hash in real-time. 

## Disclaimer
Hashunter is intended for educational and research purposes only. 
By using this tool, you agree to take full responsability for any actions you undertake. The author and contributors are not liable for any misuse or damage that may occour as a result of using this software. 
