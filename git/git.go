/*
*
https://github.com/go-git/go-git
https://git-scm.com/book/pt-br/v2/Fundamentos-de-Git-Vendo-o-histórico-de-Commits
*/
package git

import (
	"fmt"
	"github.com/inocencio/grinvon/conv"
	. "github.com/inocencio/grinvon/utils"
	"strings"
)

type CloneAltMethod int

const (
	// RemoveAndClone remove current Path directory and clone it again
	RemoveAndClone CloneAltMethod = iota
	// Pull just pull when Path directory is found with the same Branch.
	// If branches are different, the Clone alternative method switches to Checkout.
	Pull
	// Checkout the branch regardless of local branch
	Checkout
)

type CloneOptions struct {
	// Path is the directory where clone will be achieved
	Path string
	// URL is a possible remote endpoint repository to be cloned
	URL string
	// Branch indicates an alternative branch instead of default's one (main, master, or develop)
	Branch string
	// Method do some action when Path is found.
	Method CloneAltMethod
}

func IsLocalBranchEqualsRemote(path, branch string) bool {
	cmd := Cmd(path, "git", "branch", "--show-current")
	if cmd.Error != nil {
		Raise(fmt.Errorf("it was not possible to compare the branches (local and remote): " + branch + ", erro: " + cmd.Error.Error()))
	}

	return conv.IsEqualsValues(cmd.Output, branch)
}

// TODO Clone()
func Clone(dir, remote, branch string) {
	//cmd := Cmd(dir, "git", "clone", remote, "-b", branch)
	//if cmd.Error != nil {
	//	Raise(fmt.Errorf("it was not possible to clone the repository: " + cmd.Error.Error()))
	//}

	cpt := CheckPathType(dir)
	if DirType == cpt {
		if IsLocalBranchEqualsRemote(dir, branch) {

		}
	}

}

// HasCommitHash checks if a specified git commit hash is present
// in any of the branches in the git repository located in the given path.
// The function executes the `git branch --contains` command and parses its output.
//
// Parameters:
//   - path: The path to the directory containing the Git repository
//   - hash: The specified git commit hash to check for its presence in branches
//
// Returns:
//   - bool: Returns true if the specified hash is present in any branch, false otherwise.
//
// If the function encounters an execution error, it raises an error with the `Raise` function.
func HasCommitHash(path, hash string) bool {
	cmd := Cmd(path, "git", "branch", "--contains", hash)

	if cmd.Error != nil {
		Raise(fmt.Errorf("Não foi possível comparar o hash entre as branches: " + cmd.Error.Error()))
	}

	lines := strings.Split(strings.TrimSuffix(cmd.Output, "\n"), "\n")
	for _, line := range lines {
		if strings.Contains(line, "*") {
			return true
		}
	}
	return false
}

// GetLastCommitHash returns the last commit hash of the specified branch in the given path.
// If the branch does not contain any commits or an error occurred, an empty string is returned.
// The notMerged parameter is used to determine if only non-merged commits should be included in the search.
//
// Parameters:
//   - path: The path to the directory containing the Git repository
//   - branch: The branch name
//   - notMerged: Flag indicating if only non-merged commits should be considered
//
// Returns:
//   - error: Não foi encontrado a branch
//   - string: The last commit hash of the branch
func GetLastCommitHash(path, branch string, notMerged bool) (error, string) {
	if !strings.Contains(branch, "origin") {
		branch = "origin/" + branch
	}

	cmd := Cmd(path, "git", "log", branch, "--pretty=\"%H | %D | %cD | %s\"", "--abbrev-commit", "-10")
	if cmd.Error != nil {
		return fmt.Errorf("Não foi possível obter os dados da branch: " + branch), ""
	}

	for _, line := range cmd.Lines {
		if !notMerged || !strings.Contains(line, "Merge") {
			hash := strings.Split(line, " ")[0]
			hash = strings.Replace(hash, "\"", "", -1)
			return nil, hash
		}
	}

	return nil, ""
}

func FetchOrigin(path string) error {
	return Fetch(path, "origin")
}

func Fetch(path, remote string) error {
	cmd := Cmd(path, "git", "fetch", remote)
	if cmd.Error != nil {
		return fmt.Errorf("Não foi possível executar o fatch no caminho '" + path + "' para o remoto '" + remote + "'")
	}

	return nil
}
