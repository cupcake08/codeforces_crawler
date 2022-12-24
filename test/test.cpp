#include <bits/stdc++.h>
#define ll long long
#define ls u << 1
#define rs u << 1 | 1
#define mm(x) memset(x, 0, sizeof(x))
using namespace std;
int read()
{
    int a = 0;
    int f = 0;
    char p = getchar();
    while (!isdigit(p))
    {
        f |= p == '-';
        p = getchar();
    }
    while (isdigit(p))
    {
        a = (a << 3) + (a << 1) + (p ^ 48);
        p = getchar();
    }
    return f ? -a : a;
}
const int INF = 998244353;
const int P = 998244353;
const int N = 1e6 + 5;
int T;
int n, m, q;
vector<int> G1[N], G2[N];
int ff[N];
int dfn[N], low[N], cnt;
int val[N];
int sval[N];
int sum;
int t[N], top;
void tarjan(int u)
{
    t[++top] = u;
    dfn[u] = low[u] = ++cnt;
    for (auto v : G1[u])
    {
        if (!dfn[v])
        {
            tarjan(v);
            low[u] = min(low[u], low[v]);
            if (low[v] == dfn[u])
            {
                ++sum;
                for (int x = 0; x != v; --top)
                {
                    x = t[top];
                    G2[x].push_back(n + sum);
                    G2[n + sum].push_back(x);
                    // cout<<"- "<<n+sum<<" "<<x<<endl;
                }
                G2[u].push_back(n + sum);
                G2[n + sum].push_back(u);
                ff[n + sum] = u;
                // cout<<"- "<<n+sum<<" "<<u<<endl;
            }
        }
        else
            low[u] = min(low[u], dfn[v]);
    }
}
int bz[N][20];
int dep[N];
void dfs(int u, int fa)
{
    bz[u][0] = fa;
    dep[u] = dep[fa] + 1;
    for (int k = 1; k < 20; ++k)
        bz[u][k] = bz[bz[u][k - 1]][k - 1];
    for (auto v : G2[u])
    {
        if (v == fa)
            continue;
        dfs(v, u);
    }
}
int lca(int x, int y)
{
    if (dep[x] < dep[y])
        swap(x, y);
    for (int k = 19; k >= 0; --k)
        if (dep[bz[x][k]] >= dep[y])
            x = bz[x][k];
    if (x == y)
        return x;
    for (int k = 19; k >= 0; --k)
        if (bz[x][k] != bz[y][k])
            x = bz[x][k], y = bz[y][k];
    return bz[x][0];
}
void dfsv(int u, int fa, int S)
{
    S += val[u];
    sval[u] = S;
    for (auto v : G2[u])
    {
        if (v == fa)
            continue;
        dfsv(v, u, S);
    }
}
int main()
{
    n = read();
    m = read();
    for (int i = 1; i <= m; ++i)
    {
        int x = read();
        int y = read();
        G1[x].push_back(y);
        G1[y].push_back(x);
    }
    tarjan(1);
    dfs(1, 0);
    for (int u = 1; u <= n; ++u)
    {
        for (auto v : G1[u])
        {
            int L = lca(u, v);
            if (dep[u] + dep[v] == dep[L] + dep[L] + 2)
            {
                int pos = L;
                if (L == u)
                    pos = bz[v][0];
                if (L == v)
                    pos = bz[u][0];
                if (G2[pos].size() == 2)
                    continue;
                val[pos]++;
                // cout<<"--- "<<u<<" "<<v<<" "<<L<<" "<<pos<<" "<<endl;
            }
        }
    }
    dfsv(1, 0, 0);
    q = read();
    while (q--)
    {
        int x = read();
        int y = read();
        int z = lca(x, y);
        printf("%d\n", (sval[x] + sval[y] - sval[z] - sval[z] + val[z]) / 2);
    }
    return 0;
}